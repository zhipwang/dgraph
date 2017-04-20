import React from 'react';

import Editor from "../containers/Editor";

import '../assets/css/EditorPanel.css';

class EditorPanel extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      query: ''
    };
  }

  handleQueryUpdate = (val) => {
    this.setState({ query: val });
  }

  render() {
    const { onQueryRun } = this.props;

    return (
      <div className="editor-panel">
        <div className="actions">
          <ul className="action-list">
            <li className="action">
              <a
                href="#"
                className="btn btn-dgraph"
                onClick={(e) => {
                  e.preventDefault();
                  const { query } = this.state;

                  onQueryRun(query);
                }}
              >
                <i className="fa fa-play" />
              </a>
            </li>
            <li className="action">
              <a href="#" className="btn btn-dgraph">
                <i className="fa fa-close" />
              </a>
            </li>
            <li className="action">
              <a href="#" className="btn btn-dgraph">
                Share
              </a>
            </li>
          </ul>
        </div>

        <Editor
          onQueryUpdate={this.handleQueryUpdate}
        />
      </div>
    );
  }
}

export default EditorPanel;
