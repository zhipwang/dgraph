import React from 'react';

import Highlight from './Highlight';

const SessionQueryTab = ({ session }) => {
  return (
    <div>
      <Highlight>
        {session.query}
      </Highlight>
    </div>
  );
};
export default SessionQueryTab;
