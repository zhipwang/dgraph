export const RECEIVE_SESSION = 'RECEIVE_SESSION';

export function receiveSession({ id, query, response }) {
  return {
    type: RECEIVE_SESSION,
    session: {
      id,
      query,
      response
    }
  }
}
