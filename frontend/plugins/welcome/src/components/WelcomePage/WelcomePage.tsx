import React, { useState,useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import ComponanceTable from '../Table';
import Button from '@material-ui/core/Button';
import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
  Link,
} from '@backstage/core';

function WelcomePage(props: { name: React.ReactNode; }){
  const profile = { givenName: 'to System Analysis and Design 63' };

  return (
    <Page theme={pageTheme.home}>
      <Header
        title={`Welcome ${profile.givenName || 'to Backstage'}`}
        subtitle="Some quick intro and links."
      ></Header>
      <Content>
        <ContentHeader title="Application CRUD">
  <p>{props.name}</p>
          <Link component={RouterLink} to="/booking">
            <Button variant="contained" color="primary">
              Booking
           </Button>
          </Link>
        </ContentHeader>
        <ComponanceTable></ComponanceTable>
      </Content>
    </Page>
  );
};
export default WelcomePage;