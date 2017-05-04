import React from "react";
import { connect } from "react-redux";

import Sidebar from '../components/Sidebar';
import EditorPanel from '../components/EditorPanel';
import FrameList from '../components/FrameList';
import { runQuery, runQueryByShareId } from "../actions";
import { discardFrame } from '../actions/frames';
import { readCookie, eraseCookie } from './Helpers';

import "../assets/css/App.css";

class App extends React.Component {
  componentDidMount = () => {
    const { handleRunQuery, match } = this.props;

    const { shareId } = match.params;
    if (shareId) {
      this.onRunSharedQuery(shareId);
    }

    // If playQuery cookie is set, run the query and erase the cookie
    // The cookie is used to communicate the query string between docs and play
    const playQuery = readCookie('playQuery');
    if (playQuery) {
      const queryString = decodeURI(playQuery);
      handleRunQuery(queryString).then(() => {
        eraseCookie('playQuery', { crossDomain: true });
      });
    }
  };

  onRunSharedQuery(shareId) {
    const { handleRunSharedQuery } = this.props;

    handleRunSharedQuery(shareId).catch(e => {
      console.log(e);
    })
  }

  render = () => {
    const { handleRunQuery, handleDiscardFrame, frames } = this.props;

    return (
      <div className="app-layout">
        <Sidebar />
        <div className="main-content">
          <div className="container-fluid">
            <div className="row justify-content-md-center">
              <div className="col-sm-12">
                <EditorPanel
                  onRunQuery={handleRunQuery}
                />
              </div>

              <div className="col-sm-12">
                <FrameList
                  frames={frames}
                  onDiscardFrame={handleDiscardFrame}
                />
              </div>
            </div>
          </div>
        </div>
      </div>
    );
  };
}

const mapStateToProps = state => ({
  frames: state.frames.items
});

const mapDispatchToProps = dispatch => ({
  handleRunQuery(query, done = () => {}) {
    return dispatch(runQuery(query))
      .then(done);
  },
  handleRunSharedQuery(shareId) {
    return dispatch(runQueryByShareId(shareId));
  },
  handleDiscardFrame(frameID) {
    dispatch(discardFrame(frameID));
  }
});

export default connect(mapStateToProps, mapDispatchToProps)(App);
