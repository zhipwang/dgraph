import React from 'react';
import classnames from 'classnames';

import { collapseQuery } from '../containers/Helpers';

const QueryPreview = ({ query, onSelectQuery }) => {
  return (
    <div
      className={classnames('query-row')}
      onClick={(e) => {
        e.preventDefault();
        onSelectQuery(query);

        // Scroll to top
        // IDEA: This breaks encapsulation. Is there a better way?
        document.querySelector('.main-content').scrollTop = 0;
      }}
    >
      <div>
        <i className="fa fa-search query-prompt" /> <span className="preview">{collapseQuery(query)}</span>
      </div>
    </div>
  );
};

export default QueryPreview;
