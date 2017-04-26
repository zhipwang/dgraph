import React from 'react';
import classnames from 'classnames';

import Highlight from './Highlight';

const SessionJSONTab = ({ session, active }) => {
  return (
    <div className={classnames('content-container')}>
      <Highlight preClass="content">
        {JSON.stringify(session.response.data, null, 2)}
      </Highlight>
    </div>
  );
};
export default SessionJSONTab;
