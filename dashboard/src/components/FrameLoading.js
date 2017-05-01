import React from 'react';

import loader from '../assets/images/loader.svg';


const FrameLoading = ({ message }) => {
  return (
    <div className="loading">
      <img src={loader} alt="" className="loader" />
      <div className="text">Fetching result...</div>
    </div>
  );
};

export default FrameLoading;
