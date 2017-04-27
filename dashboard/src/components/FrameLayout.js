import React from 'react';

import '../assets/css/Frames.css';

class FrameLayout extends React.Component {
  handleMaximize = () => {
    //todo
  }

  handleMinimize = () => {
    //todo
  }

  render() {
    const { children } = this.props;

    return (
      <div className="frame-item">
        <div className="header">
          <div className="actions">
            <a
              href="#discard"
              className="action"
            >
              <i className="fa fa-close" />
            </a>
            <a
              href="#fullscreen"
              className="action"
              onClick={this.handleMinimize}
            >
              <i className="fa fa-minus" />
            </a>
            <a
              href="#fullscreen"
              className="action"
              onClick={this.handleMaximize}
            >
              <i className="fa fa-expand" />
            </a>
            <a
              href="#share"
              className="action"
            >
              <i className="fa fa-share-alt" />
            </a>
          </div>
        </div>

        {children}
      </div>
    );
  }
}

export default FrameLayout;
