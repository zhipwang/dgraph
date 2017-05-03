import SHA256 from "crypto-js/sha256";

import {
    timeout,
    checkStatus,
    isNotEmpty,
    processGraph,
    getEndpoint,
    makeFrame
} from "../containers/Helpers";
import {
  FRAME_TYPE_SESSION, FRAME_TYPE_SUCCESS, FRAME_TYPE_LOADING, FRAME_TYPE_ERROR
} from '../lib/const';

import { receiveFrame, updateFrame } from './frames';

// executeQueryAndUpdateFrame fetches the query response from the server
// and updates the frame
function executeQueryAndUpdateFrame(dispatch, { frameId, query }) {
  const endpoint = getEndpoint('query', { debug: true });

  return timeout(
    60000,
    fetch(endpoint, {
      method: "POST",
      mode: "cors",
      headers: {
        "Content-Type": "text/plain"
      },
      body: query
    })
  )
    .then(checkStatus)
    .then(response => response.json())
    .then((result) => {
      if (result.code !== undefined && result.message !== undefined) {
        // This is the case in which user sends a mutation.
        // We display the response from server.
        let frameType;
        if (result.code.startsWith("Error")) {
          frameType = FRAME_TYPE_ERROR;
        } else {
          frameType = FRAME_TYPE_SUCCESS;
        }

        dispatch(updateFrame({
          id: frameId,
          type: frameType,
          data: {
            query,
            message: result.message,
            response: JSON.stringify(result)
          }
        }));
      } else if (isNotEmpty(result)) {
        const { nodes, edges, labels, nodesIndex, edgesIndex } =
          processGraph(result, false, query, '');

        dispatch(updateFrame({
          id: frameId,
          type: FRAME_TYPE_SESSION,
          data: {
            query,
            response: {
              plotAxis: labels,
              allNodes: nodes,
              allEdges: edges,
              numNodes: nodes.length,
              numEdges: edges.length,
              nodes: nodes.slice(0, nodesIndex),
              edges: edges.slice(0, edgesIndex),
              treeView: false,
              data: result
            }
          }
        }));
    } else {
      dispatch(updateFrame({
        id: frameId,
        type: FRAME_TYPE_SUCCESS,
        data: {
          query,
          message: 'Your query did not return any results',
          response: JSON.stringify(result)
        }
      }));
    }
  })
  .catch((error) => {
    console.log(error.stack);
    dispatch(updateFrame({
      id: frameId,
      type: FRAME_TYPE_ERROR,
      data: {
        query,
        message: error.message,
        response: JSON.stringify(error)
      }
    }));
  })
}

/**
 * runQuery runs the query and displays the appropriate result in a frame
 * @params query {String}
 * @params [frameId] {String}
 *
 * If frameId is not given, It inserts a new frame. Otherwise, it updates the
 * frame with the id equal to frameId
 *
 */
export const runQuery = (query, frameId) => {
  return dispatch => {
    // Either insert a new frame or update
    let targetFrameId;
    if (frameId) {
      dispatch(updateFrame({
        id: frameId,
        type: FRAME_TYPE_LOADING,
        data: {}
      }));
      targetFrameId = frameId;
    } else {
      const frame = makeFrame({
        type: FRAME_TYPE_LOADING,
        data: {}
      });
      dispatch(receiveFrame(frame));
      targetFrameId = frame.id;
    }

    return executeQueryAndUpdateFrame(dispatch, {
      frameId: targetFrameId,
      query
    });
  };
};

export const addScratchpadEntry = entry => ({
    type: "ADD_SCRATCHPAD_ENTRY",
    ...entry
});

export const deleteScratchpadEntries = () => ({
    type: "DELETE_SCRATCHPAD_ENTRIES"
});

// createShare persists the queryText in the database
const createShare = (queryText) => {
  const stringifiedQuery = encodeURI(queryText);

  return fetch(getEndpoint('share'), {
    method: "POST",
    mode: "cors",
    headers: {
      Accept: "application/json",
      "Content-Type": "text/plain"
    },
    body: stringifiedQuery
  })
    .then(checkStatus)
    .then(response => response.json())
    .then((result) => {
      if (result.uids && result.uids.share) {
        return result.uids.share;
      }
    })
};

/**
 * getShareId gets the id used to share a query either by fetching one if one
 * exists, or persisting the queryText into the database.
 *
 * @params queryText {String} - A raw query text as entered by the user
 * @returns {Promise}
 */
export const getShareId = (queryText) => {
  const encodedQuery = encodeURI(queryText);
  const queryHash = SHA256(encodedQuery).toString();
  const checkQuery = `
{
  query(func:eq(_share_hash_, ${queryHash})) {
      _uid_
      _share_
  }
}`;

  return timeout(
    6000,
    fetch(getEndpoint('query'), {
      method: "POST",
      mode: "cors",
      headers: {
        Accept: "application/json",
        "Content-Type": "text/plain"
      },
      body: checkQuery
    })
      .then(checkStatus)
      .then(response => response.json())
      .then((result) => {
        const matchingQueries = result.query;

        // If no match, store the query
        if (!matchingQueries) {
          return createShare(queryText);
        }

        if (matchingQueries.length === 1) {
          return result.query[0]._uid_;
        }

        // If more than one result, we have a hash collision. Break it.
        for (let i = 0; i < matchingQueries.length; i++) {
          const q = matchingQueries[i];
          if (`"${q._share_}"` === encodedQuery) {
            return q._uid_;
          }
        }
      })
  );
};

// export const getQuery = shareId => {
//     return dispatch => {
//         timeout(
//             6000,
//             fetch(getEndpoint('query'), {
//                 method: "POST",
//                 mode: "cors",
//                 headers: {
//                     Accept: "application/json"
//                 },
//                 body: `{
//                     query(id: ${shareId}) {
//                         _share_
//                     }
//                 }`
//             })
//                 .then(checkStatus)
//                 .then(response => response.json())
//                 .then(function(result) {
//                     if (result.query && result.query.length > 0) {
//                         dispatch(selectQuery(decodeURI(result.query[0]._share_)));
//                     } else {
//                         dispatch(queryFound(false));
//                     }
//                 })
//         ).catch(function(error) {
//             dispatch(
//                 saveErrorResponse(
//                     `Got error while getting query for id: ${shareId}, err: ` +
//                         error.message
//                 )
//             );
//         });
//     };
// };
