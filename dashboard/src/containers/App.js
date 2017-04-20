import React from "react";
import ReactDOM from 'react-dom';
import { connect } from "react-redux";
import screenfull from "screenfull";
import { Alert } from "react-bootstrap";

import Sidebar from '../components/Sidebar';
import EditorPanel from '../components/EditorPanel';

import {
  updateFullscreen,
  getQuery,
  updateInitialQuery,
  queryFound,
  initialServerState,
  selectQuery,
  runQuery
} from "../actions";

import "../assets/css/App.css";

class App extends React.Component {
  enterFullScreen = () => {
    if (!screenfull.enabled) {
      return;
    }

    const responseEl = ReactDOM.findDOMNode(this.refs.response);
    screenfull.request(responseEl);
  };

  toggleSidebarOpen = () => {
    this.setState({ sidebarOpen: !this.state.sidebarOpen });
  }

  render = () => {
    const { handleQueryRun } = this.props;

    return (
      <div className="app-layout">
        <Sidebar />
        <div className="main-content">
          <div className="container-fluid">
            <div className="row justify-content-md-center">
              <div className="col-sm-12">
                {!this.props.found &&
                  <Alert
                    ref={alert => {
                      this.alert = alert;
                    }}
                    bsStyle="danger"
                    onDismiss={() => {
                      this.props.queryFound(true);
                    }}
                  >
                    Couldn't find query with the given id.
                  </Alert>}
              </div>
              <div>
                <EditorPanel
                  onQueryRun={handleQueryRun}
                />
              </div>
            </div>
          </div>
        </div>
      </div>
    );
  };

  componentDidMount = () => {

    const { handleUpdateFullscreen } = this.props;
    this.props.initialServerState();

    const { handleQuerySelect, handleQueryRun } = this.props;

    let id = this.props.match.params.id;
    if (id !== undefined) {
      this.props.getQuery(id);
    }

    // If playQuery cookie is set, run the query and erase the cookie
    // The cookie is used to communicate the query string between docs and play
    const playQuery = readCookie('playQuery');
    if (playQuery) {
      const queryString = decodeURI(playQuery);
      handleQuerySelect(queryString);
      handleQueryRun(queryString).then(() => {
        eraseCookie('playQuery', { crossDomain: true });
      });
    }
  };

  componentWillUnmount = () => {
    const { handleUpdateFullscreen } = this.props;

    document.removeEventListener(screenfull.raw.fullscreenchange, handleUpdateFullscreen);
  }

  componentWillReceiveProps = nextProps => {
    if (!nextProps.found) {
      // Lets auto close the alert after 2 secs.
      setTimeout(
        () => {
          this.props.queryFound(true);
        },
        3000
      );
    }
  };
}

const mapStateToProps = state => ({
  found: state.share.found
});

const mapDispatchToProps = dispatch => ({
  handleUpdateFullscreen: () => {
    const fsState = screenfull.isFullscreen;
    dispatch(updateFullscreen(fsState));
  },
  getQuery: id => {
    dispatch(getQuery(id));
  },
  updateInitialQuery: () => {
    dispatch(updateInitialQuery());
  },
  queryFound: found => {
    dispatch(queryFound(found));
  },
  initialServerState: () => {
    dispatch(initialServerState());
  },
  handleQuerySelect: (queryText) => {
    dispatch(selectQuery(queryText));
  },
  handleQueryRun: (query) => {
    return dispatch(runQuery(query));
  }
});

export default connect(mapStateToProps, mapDispatchToProps)(App);
