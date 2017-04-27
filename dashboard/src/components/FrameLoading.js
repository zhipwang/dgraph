import React from 'react';

import loader from '../assets/images/loader.svg';

const FrameLoading = ({ message }) => {
  return (
    <div className="loading">
      <img src={loader} alt="loader"/>
    </div>
  );
};

export default FrameLoading;
