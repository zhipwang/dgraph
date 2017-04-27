export const RECEIVE_FRAME = 'RECEIVE_FRAME';
export const DISCARD_FRAME = 'DISCARD_FRAME';

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
