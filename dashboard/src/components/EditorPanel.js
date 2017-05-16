import React from 'react';
import classnames from 'classnames';

import Editor from "../containers/Editor";

import '../assets/css/EditorPanel.css';

class EditorPanel extends React.Component {
  render() {
    const { isQueryDirty, query, onRunQuery, onUpdateQuery, onClearQuery, saveCodeMirrorInstance } = this.props;

    return (
      <div className="editor-panel">
        <div className="header">
          <div className="status">
            <i className="fa fa-circle icon" /> <span className="text">Connected</span>
          </div>
          <div className="actions">
            <a
              href="#"
              className={classnames('action clear-btn', { actionable: isQueryDirty })}
              onClick={(e) => {
                e.preventDefault();
                if (query === '') {
                  return;
                }

                onClearQuery();
              }}
            >
              <i className="fa fa-close" />
            </a>
            <a
              href="#"
              className={classnames('action run-btn', { actionable: isQueryDirty })}
              onClick={(e) => {
                e.preventDefault();
                if (query === '') {
                  return;
                }

                onRunQuery(query);
              }}
            >
              <i className="fa fa-play"/>
            </a>
          </div>
        </div>

        <Editor
          onUpdateQuery={onUpdateQuery}
          onRunQuery={onRunQuery}
          query={query}
          saveCodeMirrorInstance={saveCodeMirrorInstance}
        />
      </div>
    );
  }
}

export default EditorPanel;
