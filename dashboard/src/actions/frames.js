export const RECEIVE_FRAME = 'RECEIVE_FRAME';

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
