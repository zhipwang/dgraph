import React from 'react';
import classnames from 'classnames';

import GraphContainer from "../containers/GraphContainer";

const SessionTreeTab = ({
  session, active, onBeforeTreeRender, onTreeRendered, onNodeSelected,
  onNodeHovered, selectedNode }) => {
  return (
    <div className={classnames('content-container', { hidden: !active})}>
      <GraphContainer
        response={session.response}
        onBeforeRender={onBeforeTreeRender}
        onRendered={onTreeRendered}
        onNodeSelected={onNodeSelected}
        onNodeHovered={onNodeHovered}
        selectedNode={selectedNode}
        treeView
      />
    </div>
  );
};

export default SessionTreeTab;
