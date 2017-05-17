import React from "react";

import SessionFooterResult from "./SessionFooterResult";
import SessionFooterProperties from "./SessionFooterProperties";

const SessionFooterDisplayConfig = () => {
  return <div>SessionFooterDisplayConfig</div>;
};

const SessionFooter = ({
  session,
  currentTab,
  graphRenderTime,
  treeRenderTime,
  hoveredNode,
  selectedNode
}) => {
  let child;
  if (selectedNode) {
    child = <SessionFooterProperties entity={selectedNode} />;
  } else if (hoveredNode) {
    child = <SessionFooterProperties entity={hoveredNode} />;
  } else {
    child = (
      <SessionFooterResult
        currentTab={currentTab}
        session={session}
        graphRenderTime={graphRenderTime}
        treeRenderTime={treeRenderTime}
      />
    );
  }

  return (
    <div className="footer">
      {child}
    </div>
  );
};
export default SessionFooter;
