import { createPlugin, createRouteRef } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import Booking from './components/Booking';
import Login from './components/Login';
export const rootRouteRef = createRouteRef({
  path: '/welcome',
  title: 'welcome',
});

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.addRoute(rootRouteRef, WelcomePage);
    router.registerRoute('/booking', Booking);
    router.registerRoute('/login', Login);
  },
});
