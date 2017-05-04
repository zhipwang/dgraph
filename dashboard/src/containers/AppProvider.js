import React from 'react';
import { Provider } from 'react-redux';
import { compose, createStore, applyMiddleware } from 'redux';
import { persistStore, autoRehydrate } from 'redux-persist';
import {
  BrowserRouter as Router,
  Route,
  browserHistory
} from 'react-router-dom';
import thunk from 'redux-thunk';
import reducer from "../reducers";

import "bootstrap/dist/css/bootstrap.css";
import "bootstrap/dist/css/bootstrap-theme.css";

const middleware = [thunk];
const composeEnhancers = window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose;
const store = createStore(
  reducer,
  undefined,
  composeEnhancers(applyMiddleware(...middleware), autoRehydrate())
);

export default class AppProvider extends React.Component {
  constructor() {
    super();
    this.state = {
      rehydrated: false
    };
  }

  componentWillMount() {
    // begin periodically persisting the store
    persistStore(store, { whitelist: ["frames"] }, () => {
      this.setState({ rehydrated: true });
    });
  }

  render() {
    const {component} = this.props;
    const { rehydrated } = this.state;

    if (!rehydrated) {
      return (
        <div>
          Loading
        </div>
      )
    }

    return (
      <Provider store={store}>
        <Router history={browserHistory}>
          <Route path="/:shareId?" component={component} />
        </Router>
      </Provider>
    )
  }
}
