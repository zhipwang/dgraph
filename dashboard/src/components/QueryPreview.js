import React from 'react';
import classnames from 'classnames';

import { collapseQuery } from '../containers/Helpers';

const QueryPreview = ({ query, editingQuery, onToggleEditingQuery }) => {
  return (
    <div
      className={classnames('query-row', { editing: editingQuery })}
      onClick={(e) => {
        e.preventDefault();
        if (editingQuery) {
          return;
        }

        onToggleEditingQuery();
      }}
    >
      {editingQuery ? null :
        <div>
          <i className="fa fa-search query-prompt" /> <span className="preview">{collapseQuery(query)}</span>
        </div>
      }
    </div>
  );
};
export default QueryPreview;
