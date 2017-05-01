import React from 'react';
import ReactDOM from 'react-dom';
import screenfull from 'screenfull';
import classnames from 'classnames';

import FrameHeader from './FrameHeader';
import { FRAME_TYPE_SESSION } from '../lib/const';
import { getShareId } from '../actions';

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
      const shareUrlEl = ReactDOM.findDOMNode(this.shareURLEl);
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
      const shareUrlEl = ReactDOM.findDOMNode(this.shareURLEl);

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

  // saveShareURLRef saves the reference to the share url input as an instance
  // property of this component
  saveShareURLRef = (el) => {
    this.shareURLEl = el;
  }

  render() {
    const { children, onDiscardFrame, frame } = this.props;
    const { isFullscreen, shareId, shareHidden } = this.state;

    return (
      <div
        className={classnames('frame-item', { fullscreen: isFullscreen })}
        ref="frame"
      >
        <FrameHeader
          shareId={shareId}
          onToggleFullscreen={this.handleToggleFullscreen}
          onDiscardFrame={onDiscardFrame}
          onShare={this.handleShare}
          shareHidden={shareHidden}
          frame={frame}
          isFullscreen={isFullscreen}
          saveShareURLRef={this.saveShareURLRef}
        />

        {children}
      </div>
    );
  }
}

export default FrameLayout;
