import React from 'react';

import Highlight from './Highlight';

const SessionCodeTab = ({ session, active }) => {
  return (
    <div className="content-container">
      <div className="code-container">
        <div className="code-header">
          <span className="label label-info">
            Query
          </span>
        </div>
        <Highlight preClass="content">
          {session.query}
        </Highlight>
      </div>

      <div className="code-container">
        <div className="code-header">
          <span className="label label-info">
            Response
          </span>
        </div>
        <Highlight preClass="content">
          {JSON.stringify(session.response.data, null, 2)}
        </Highlight>
      </div>
    </div>
  );
};
export default SessionCodeTab;
