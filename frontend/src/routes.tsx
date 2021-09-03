import React from 'react';
import { BrowserRouter, Switch, Route } from 'react-router-dom'

import CreateOrphanage from './pages/CreateOrphanage';
import Landing from './pages/Landing';
import CreateUser from './pages/CreateUser';
import Orphanage from './pages/Orphanage';

export default function Routes() {
  return (
    <BrowserRouter>
      <Switch>
        <Route path="/" exact component={Landing} />
        <Route path="/appointment/create" component={CreateOrphanage} />
        <Route path="/user/create" component={CreateUser} />
        <Route path="/user/:id" component={Orphanage} />
        <Route path="/appointment/:id" component={Orphanage} />
      </Switch>
    </BrowserRouter>
  );
}
