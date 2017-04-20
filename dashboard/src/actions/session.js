export const RECEIVE_SESSION = 'RECEIVE_SESSION';

export function receiveSession({ query, result }) {
  return {
    type: RECEIVE_SESSION,
    session: {
      query,
      result
    }
  }
}
