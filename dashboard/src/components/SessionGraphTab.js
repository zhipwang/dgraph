import React from "react";

import GraphContainer from "../containers/GraphContainer";

const SessionGraphTab = ({
  session,
  active,
  onBeforeGraphRender,
  onGraphRendered,
  onNodeSelected,
  onNodeHovered,
  selectedNode,
  hoveredNode,
  nodesDataset,
  edgesDataset,
  labelRegex,
  applyLabels
}) => {
  return (
    <div className="content-container">
      <GraphContainer
        response={session.response}
        onBeforeRender={onBeforeGraphRender}
        onRendered={onGraphRendered}
        onNodeSelected={onNodeSelected}
        onNodeHovered={onNodeHovered}
        selectedNode={selectedNode}
        hoveredNode={hoveredNode}
        nodesDataset={nodesDataset}
        edgesDataset={edgesDataset}
        labelRegex={labelRegex}
        applyLabels={applyLabels}
      />
    </div>
  );
};

export default SessionGraphTab;
