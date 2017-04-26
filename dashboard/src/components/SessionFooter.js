import React from 'react';

import Properties from '../components/Properties';
import { humanizeTime } from '../containers/Helpers';

const MetaInfo = ({ session, currentTab, graphRenderTime, treeRenderTime }) => {
  return (
    <ul className="stats">
      {session.response.data.server_latency ?
        <li className="stat">Server latency: <span className="value">{session.response.data.server_latency.total}</span></li> : null}
      {graphRenderTime && currentTab === 'graph' ?
        <li className="stat">Rendering latency: <span className="value">{humanizeTime(graphRenderTime)}</span></li> : null}
      {treeRenderTime && currentTab === 'tree' ?
        <li className="stat">Rendering latency: <span className="value">{humanizeTime(treeRenderTime)}</span></li> : null}
      <li className="stat">
        Nodes: <span className="value">{session.response.numNodes}</span>
      </li>
      <li className="stat">
        Edges: <span className="value">{session.response.numEdges}</span>
      </li>
    </ul>
  );
}

const NodeInfo = ({ node }) => {
  return <Properties node={node} />;
}

const SessionFooter = ({
  session, currentTab, selectedNode, hoveredNode, graphRenderTime, treeRenderTime
}) => {
  const nodeFocused = selectedNode || hoveredNode;

  return (
    <div className="footer">
      {nodeFocused ?
        <NodeInfo
          node={nodeFocused}
        /> :
        <MetaInfo
          session={session}
          currentTab={currentTab}
          graphRenderTime={graphRenderTime}
          treeRenderTime={treeRenderTime}
        />}
    </div>
  );
};
export default SessionFooter;
