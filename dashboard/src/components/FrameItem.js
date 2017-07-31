import React from "react";

import {
  FRAME_TYPE_SESSION,
  FRAME_TYPE_SUCCESS,
  FRAME_TYPE_ERROR,
  FRAME_TYPE_LOADING
} from "../lib/const";
import FrameLayout from "./FrameLayout";
import FrameSession from "./FrameSession";
import FrameError from "./FrameError";
import FrameSuccess from "./FrameSuccess";
import FrameLoading from "./FrameLoading";

// getFrameContent returns React Component for a given frame depending on its type
function getFrameContent(frame, response) {
  if (frame.type === FRAME_TYPE_SESSION) {
    return <FrameSuccess data={frame.data} response={response} />;
    // return <FrameSession session={frame.data} response={response} />;
  } else if (frame.type === FRAME_TYPE_ERROR) {
    return <FrameError data={frame.data} />;
  } else if (frame.type === FRAME_TYPE_SUCCESS) {
    return <FrameSuccess data={frame.data} response={response} />;
  } else if (frame.type === FRAME_TYPE_LOADING) {
    return <FrameLoading />;
  }

  return <FrameError message={`Unknown frame type: ${frame.type}`} />;
}

const FrameItem = ({
  frame,
  response,
  onDiscardFrame,
  onSelectQuery,
  collapseAllFrames
}) => {
  const content = getFrameContent(frame, response);

  return (
    <FrameLayout
      frame={frame}
      response={response}
      onDiscardFrame={onDiscardFrame}
      onSelectQuery={onSelectQuery}
      collapseAllFrames={collapseAllFrames}
    >
      {content}
    </FrameLayout>
  );
};

export default FrameItem;
