import React from 'react';
import FrameItem from './FrameItem';
import CSSTransitionGroup from 'react-transition-group/CSSTransitionGroup'

import '../assets/css/Frames.css';

const FrameList = ({ frames, onDiscardFrame }) => {
  return (
    <ul className="frame-list">
      <CSSTransitionGroup
        transitionName="frame-item"
        transitionEnterTimeout={300}
        transitionLeaveTimeout={300}
      >
        {
          frames.map((frame) => {
            return (
              <FrameItem
                key={frame.id}
                frame={frame}
                onDiscardFrame={onDiscardFrame}
              />
            )
          })
        }
      </CSSTransitionGroup>
    </ul>
  );
};

export default FrameList;
