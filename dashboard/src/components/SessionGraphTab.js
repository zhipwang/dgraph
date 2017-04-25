import React from 'react';
import classnames from 'classnames';

import GraphContainer from "../containers/GraphContainer";

const SessionGraphTab = ({
  session, active, onBeforeGraphRender, onGraphRendered, onNodeSelected, currentNode
}) => {
  return (
    <div className={classnames('content-container', { hidden: !active})}>
      <GraphContainer
        response={session.response}
        onBeforeRender={onBeforeGraphRender}
        onRendered={onGraphRendered}
        onNodeSelected={onNodeSelected}
        currentNode={currentNode}
      />
    </div>
  );
};

export default SessionGraphTab;
