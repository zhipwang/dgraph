import React from 'react';

import GraphContainer from "../containers/GraphContainer";

const SessionGraphTab = ({ session }) => {
  return (
    <div>
      <GraphContainer
        response={session.response}
      />
    </div>
  );
};

export default SessionGraphTab;
