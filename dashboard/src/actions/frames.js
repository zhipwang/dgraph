export const RECEIVE_FRAME = 'frames/RECEIVE_FRAME';
export const DISCARD_FRAME = 'frames/DISCARD_FRAME';
export const UPDATE_FRAME = 'frames/UPDATE_FRAME';

export function receiveFrame({ id, type, data }) {
  return {
    type: RECEIVE_FRAME,
    frame: {
      id,
      type,
      data
    }
  }
}

export function discardFrame(frameID) {
  return {
    type: DISCARD_FRAME,
    frameID
  };
}

export function updateFrame({ id, type, data }) {
  return {
    type: UPDATE_FRAME,
    id,
    frame: {
      type,
      data
    }
  }
}
