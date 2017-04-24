export const RECEIVE_SESSION = 'RECEIVE_SESSION';

export function receiveSession({ query, response }) {
  return {
    type: RECEIVE_SESSION,
    session: {
      query,
      response
    }
  }
}
