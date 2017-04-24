import React from 'react';
import SessionItem from './SessionItem';

import '../assets/css/SessionList.css';

const SessionList = ({ sessions }) => {
  return (
    <ul className="session-list">
      {
        sessions.map((session, idx) => {
          return (
            <SessionItem
              key={idx}
              session={session}
            />
          )
        })
      }
    </ul>
  );
};

export default SessionList;
