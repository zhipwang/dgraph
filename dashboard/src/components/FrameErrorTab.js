import React from 'react';

import Highlight from './Highlight';

const FrameErrorTab = ({ message }) => {
  return (
    <div className="content-container">
      <div className="text-content">
        {message}
      </div>
    </div>
  );
};
export default FrameErrorTab;
