import React from 'react';
import classnames from 'classnames';

import Editor from "../containers/Editor";

import '../assets/css/EditorPanel.css';

class EditorPanel extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      query: '',
      dirty: false
    };
  }

  handleQueryUpdate = (val) => {
    const dirty = val.trim() !== '';

    this.setState({ query: val, dirty });
  }

  handleRunQuery = (query) => {
    const { onRunQuery } = this.props;

    // TODO: setState should be a callback to onRunQuery
    this.setState({ dirty: false });
    onRunQuery(query);
  }

  render() {
    const { dirty } = this.state;

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
                const { query } = this.state;

                this.handleRunQuery(query);
              }}
            >
              <i className={classnames('fa fa-play', { dirty })} id="run-btn" />
            </a>

            <a href="#" className="action">
              <i className="fa fa-share-alt" />
            </a>
          </div>
        </div>

        <Editor
          onUpdateQuery={this.handleQueryUpdate}
          onRunQuery={this.handleRunQuery}
        />
      </div>
    );
  }
}

export default EditorPanel;
