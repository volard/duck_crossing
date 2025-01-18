import React from 'react';
import ReactDOM from 'react-dom/client';
import { App } from './App';
import { StoreContext, createStore } from './store';

const store = createStore();

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <StoreContext.Provider value={store}>
      <App />
    </StoreContext.Provider>
  </React.StrictMode>
);

