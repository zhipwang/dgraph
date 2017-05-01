import React from 'react';

import {
  FRAME_TYPE_SESSION, FRAME_TYPE_SYSTEM, FRAME_TYPE_ERROR, FRAME_TYPE_LOADING
} from '../lib/const';
import FrameLayout from './FrameLayout';
import FrameSession from './FrameSession';
import FrameError from './FrameError';
import FrameSystem from './FrameSystem';
import FrameLoading from './FrameLoading';


// getFrameContent returns React Component for a given frame depending on its type
function getFrameContent(frame) {
  if (frame.type === FRAME_TYPE_SESSION) {
    return (
      <FrameSession
        session={frame.data}
      />
    );
  } else if (frame.type === FRAME_TYPE_ERROR) {
    return (
      <FrameError
        data={frame.data}
      />
    )
  } else if (frame.type === FRAME_TYPE_SYSTEM) {
    return (
      <FrameSystem
        message={frame.data.message}
      />
    )
  } else if (frame.type === FRAME_TYPE_LOADING) {
    return (
      <FrameLoading />
    );
  }

  return (
    <FrameError
      message={`Unknown frame type: ${frame.type}`}
    />
  );
}

const FrameItem = ({ frame, onDiscardFrame }) => {
  const content = getFrameContent(frame);

  return (
    <FrameLayout
      frame={frame}
      onDiscardFrame={onDiscardFrame}
    >
      {content}
    </FrameLayout>
  );
};

export default FrameItem;
