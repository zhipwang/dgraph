import React from 'react';
import FrameItem from './FrameItem';
import CSSTransitionGroup from 'react-transition-group/CSSTransitionGroup'

import '../assets/css/SessionList.css';

const FrameList = ({ frames }) => {
  return (
    <ul className="frame-list">
      <CSSTransitionGroup
        transitionName="session-item"
        transitionEnterTimeout={800}
        transitionLeaveTimeout={300}
      >
        {
          frames.map((frame) => {
            return (
              <FrameItem
                key={frame.id}
                frame={frame}
              />
            )
          })
        }
      </CSSTransitionGroup>
    </ul>
  );
};

export default FrameList;
