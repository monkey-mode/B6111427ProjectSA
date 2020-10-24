import React, { useState, useEffect,Component} from 'react';
import { Link as RouterLink } from 'react-router-dom';
import WelcomePage from '../WelcomePage';
import {
  Content, Header, Page, pageTheme, ContentHeader,
} from '@backstage/core';
import {
  TextField, Button, withStyles, makeStyles,
  Theme, FormControl, InputLabel, MenuItem,
  Typography, Select, createStyles,
} from '@material-ui/core';

import { Alert, Autocomplete } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';
import {
  EntUser,
} from '../../api/models';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent: 'center',
    },
    margin: {
      margin: theme.spacing(3),
    },
    withoutLabel: {
      marginTop: theme.spacing(3),
    },
    textField: {
      width: '25ch',
    },
  }),
);

export default function Loglin(){
  const profile = { givenName: 'ยินดีต้อนรับสู่หน้า Login ระบบ VideoOnDemand' };
  const classes = useStyles();
  const api = new DefaultApi();
  const [loading, setLoading] = useState(true);
  const [users, setUsers] = useState<EntUser[]>(Array);
  useEffect(() => {
    const getUsers = async () => {
      const res = await api.listUser({ limit: 10, offset: 0 });
      setLoading(false);
      setUsers(res);
    };
    getUsers();
    
  }, [loading]);
  
  const UserIDhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setUserID(event.target.value as number);
  };

  const [userID, setUserID] = useState(Number);
  const [user, setUser] = useState(String);
  return (
    
    <Page theme={pageTheme.home}>
      <Header
        title={`${profile.givenName}`}
      ></Header>
      <Content>

        <ContentHeader title="กรุณาเลือกทำการใส่ Email"></ContentHeader>
        <div className={classes.root}>
          <form noValidate autoComplete="off">
            <table>
              <tr><td width="100">อีเมลผู้ใช้</td><td>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <InputLabel id="email-label"></InputLabel>
                <Select
                  labelId="email-label"
                  id="email"
                  value={userID}

                  onChange={UserIDhandleChange}
                  style={{ width: 400 }}
                >
                  {users.map((item: EntUser) => (
                    <MenuItem value={item.id}>{item.uSEREMAIL}</MenuItem>
                  ))}
                </Select>
              </FormControl>
              </td>
              </tr>
              <tr><td>ชื่อผู้ใช้</td><td>
              <FormControl
                disabled
                className={classes.margin}
                variant="outlined"
              >
                <InputLabel id="user-label"></InputLabel>
                <Select
                  labelId="user-label"
                  id="user"
                  value={userID}
                  onChange={UserIDhandleChange}
                  style={{ width: 400 }}
                >
                  {users.map((item: EntUser) => (
                    <MenuItem value={item.id}>{item.uSERNAME}</MenuItem>
                  ))}
                </Select>
              </FormControl>
              </td>
              </tr>
            </table>
            <table align="center">
            <div>
              <Button
              onClick={() => {
                
              }}
                style={{ marginLeft: 20 }}
                component={RouterLink}
                to="/welcome"
                variant="contained"
                color="primary"
                
              >
                [Login]
           </Button>
           </div>
           </table>
          </form>
        </div>
        
      </Content>
    </Page>
  );
}