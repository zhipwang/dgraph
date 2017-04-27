import React from 'react';
import ReactDOM from 'react-dom';
import screenfull from 'screenfull';
import classnames from 'classnames';

import { FRAME_TYPE_SESSION } from '../lib/const';
import { getShareId } from '../actions';
import { getShareURL } from '../containers/Helpers';

class FrameLayout extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      isFullscreen: false,
      shareId: '',
      shareHidden: false
    };
  }

  componentDidUpdate(prevProps, prevState) {
    // If shareId was fetched, select the share url input
    if (prevState.shareId !== this.state.shareId && this.state.shareId !== '') {
      const shareUrlEl = ReactDOM.findDOMNode(this.refs.shareUrl);
      shareUrlEl.select();
    }
  }

  handleToggleFullscreen = () => {
    if (!screenfull.enabled) {
      return;
    }

    const { isFullscreen } = this.state;

    if (isFullscreen) {
      screenfull.exit();
      this.setState({ isFullscreen: false });
    } else {
      const frameEl = ReactDOM.findDOMNode(this.refs.frame);
      screenfull.request(frameEl);

      // If fullscreen request was successful, set state
      if (screenfull.isFullscreen) {
        this.setState({ isFullscreen: true });
      }
    }
  }

  handleShare = () => {
    const { frame } = this.props;
    const { shareId } = this.state;

    // if shareId is already set, simply toggle the hidden state
    if (shareId) {
      const shareUrlEl = ReactDOM.findDOMNode(this.refs.shareUrl);

      this.setState({ shareHidden: !this.state.shareHidden });
      shareUrlEl.select();
      return;
    }

    if (frame.type !== FRAME_TYPE_SESSION) {
      return;
    }

    const { query } = frame.data;
    getShareId(query)
      .then(shareId => {
        this.setState({ shareId });
      })
      .catch(err => {
        console.log('error while getting share id', err);
      })
  }

  render() {
    const { children, onDiscardFrame, frame } = this.props;
    const { isFullscreen, shareId, shareHidden } = this.state;
    const shareURLValue = shareId ? getShareURL(shareId) : '';

    return (
      <div
        className={classnames('frame-item', { fullscreen: isFullscreen })}
        ref="frame"
      >
        <div className="header">
          <div className="actions">
            <a
              href="#share"
              className="action"
              onClick={this.handleShare}
            >
              <i className="fa fa-share-alt" />
            </a>
            <input
              type="text"
              value={shareURLValue}
              className={classnames('share-url-holder', { shared: Boolean(shareId) && !shareHidden })}
              ref="shareUrl"
              onClick={(e) => {
                e.target.select();
              }}
              onKeyUp={(e) => {
                e.target.select();
              }}
            />

            <a
              href="#fullscreen"
              className="action"
              onClick={this.handleToggleFullscreen}
            >
              {isFullscreen ?
                <i className="fa fa-compress" /> :
                <i className="fa fa-expand" />}

            </a>

            {!isFullscreen ?
              <a
                href="#discard"
                className="action"
                onClick={(e) => {
                  e.preventDefault();
                  onDiscardFrame(frame.id)
                }}
              >
                <i className="fa fa-close" />
              </a> : null}
          </div>
        </div>

        {children}
      </div>
    );
  }
}

export default FrameLayout;
