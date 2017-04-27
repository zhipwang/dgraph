import React from 'react';
import ReactDOM from 'react-dom';
import screenfull from "screenfull";

class FrameLayout extends React.Component {
  handleMaximize = () => {
    if (!screenfull.enabled) {
      return;
    }

    const frameEl = ReactDOM.findDOMNode(this.refs.frame);
    screenfull.request(frameEl);
  }

  render() {
    const { children, onDiscardFrame, frame } = this.props;

    return (
      <div className="frame-item" ref="frame">
        <div className="header">
          <div className="actions">
            <a
              href="#discard"
              className="action"
              onClick={(e) => {
                e.preventDefault();
                onDiscardFrame(frame.id)
              }}
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
