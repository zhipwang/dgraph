import React from 'react';
import SessionItem from './SessionItem';

const SessionList = ({ sessions }) => {
  return (
    sessions.map((session, idx) => {
      return (
        <SessionItem
          key={idx}
          session={session}
        />
      )
    })
  );
};
export default SessionList;
