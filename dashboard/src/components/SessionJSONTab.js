import React from 'react';

import Highlight from './Highlight';

const SessionJSONTab = ({ session }) => {
  return (
    <div>
      <Highlight preClass="json-response">
        {JSON.stringify(session.response.data, null, 2)}
      </Highlight>
    </div>
  );
};
export default SessionJSONTab;
