import React from 'react';
import classnames from 'classnames';

import Highlight from './Highlight';

const SessionQueryTab = ({ session, active }) => {
  return (
    <div className={classnames('content-container', { hidden: !active})}>
      <Highlight preClass="content">
        {session.query}
      </Highlight>
    </div>
  );
};
export default SessionQueryTab;
