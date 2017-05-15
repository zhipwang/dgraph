import React from 'react';
import classnames from 'classnames';

import Editor from "../containers/Editor";

import '../assets/css/EditorPanel.css';

class EditorPanel extends React.Component {
  render() {
    const { isQueryDirty, query, onRunQuery, onUpdateQuery } = this.props;

    return (
      <div className="editor-panel">
        <div className="header">
          <div className="status">
            <i className="fa fa-circle icon" /> <span className="text">Connected</span>
          </div>
          <div className="actions">
            <a
              href="#"
              className="action"
              onClick={(e) => {
                e.preventDefault();

                onRunQuery(query);
              }}
            >
              <i className={classnames('fa fa-play', { dirty: isQueryDirty })} id="run-btn" />
            </a>
          </div>
        </div>

        <Editor
          onUpdateQuery={onUpdateQuery}
          onRunQuery={onRunQuery}
          query={query}
        />
      </div>
    );
  }
}

export default EditorPanel;
