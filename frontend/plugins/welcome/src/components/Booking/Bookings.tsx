import React,{ useState, useEffect }  from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
  Content,Header,Page,pageTheme,ContentHeader,
} from '@backstage/core';
import {
  Table,TableBody,TableCell,TableRow,Typography,
  TextField,Button,withStyles,makeStyles,
  Theme,FormControl,InputLabel,MenuItem,
  FormHelperText,Select,createStyles,
} from '@material-ui/core';

import { Alert } from '@material-ui/lab';

import { DefaultApi } from '../../api/apis';
import { 
  EntClientEntity,
  EntBookingtype,
  EntBooking,
  EntUser,
} from '../../api/models/';

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

const username = { givenuser: 'Suphachai Phetthamrong' };

export default function Create() {
  const profile = { givenName: 'ยินดีต้อนรับสู่ ระบบ VideoOnDemand' };
  const classes = useStyles();
  const api = new DefaultApi();
  const [loading, setLoading] = useState(true);

  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  const [clients, setClients] = useState<EntClientEntity[]>(Array);
  const [users, setUsers] = useState<EntUser[]>(Array);
  const [bookingtypes, setBookingtypes] = useState<EntBookingtype[]>(Array);
useEffect(() => {
    const getCliententity = async () => {
        const res = await api.listCliententity({ limit: 10, offset: 0 });
        setLoading(false);
        setClients(res);
    };
    getCliententity();

    const getUsers = async () => {
        const res = await api.listUser({ limit: 10, offset: 0 });
        setLoading(false);
        setUsers(res);
    };
    getUsers();

    const getBookingtype = async () => {
        const res = await api.listBookingtype({ limit: 10, offset: 0 });
        setLoading(false);
        setBookingtypes(res);
    };
    getBookingtype();

    
  
    
}, [loading]);

const BookingDatehandleChange = (event: any) => {
  setBookingdate(event.target.value as string);
};

const BookingIDhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
  setBookingtypeID(event.target.value as number);
};

const clientIDhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
  setClientID(event.target.value as number);
};

const UserIDhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
  setUserID(event.target.value as number);
};


const [bookingdate, setBookingdate] = useState(String);
const [timeleft, setTimeleft] = useState(String);
const [bookingtypeID, setBookingtypeID] = useState(Number);
const [clientID, setClientID] = useState(Number);
const [userID, setUserID] = useState(Number);

const CreateBooking = async () => {
  setTimeleft("03:00:00")
  const booking = {
    bookingDate : bookingdate + ":00+07:00",
    timeLeft : timeleft,
    bookingtype : bookingtypeID,
    client : clientID,
    user : userID,
  };
  console.log(booking);
  const res: any = await api.createBooking({ booking : booking });
  setStatus(true);
    if (res.id != '') {
      setAlert(true);
    } else {
      setAlert(false);
    }
    const timer = setTimeout(() => {
      setStatus(false);
    }, 1000);
};

return (
  <Page theme={pageTheme.home}>
    <Header
      title={`${profile.givenName}`}
    //subtitle="Some quick intro and links."
    ></Header>
    <Content>
      <ContentHeader title="เพิ่มข้อมูลการจอง">
        <Typography align="left" style={{ marginRight: 16, color: "#00eeff" }}>
          {username.givenuser}
        </Typography>
        <div>
          <Button variant="contained" color="primary">
            ออกจากระบบ
       </Button>
        </div>
        {status ? (
          <div>
            {alert ? (
              <Alert severity="success">
                This is a success alert — check it out!
              </Alert>
            ) : (
                <Alert severity="warning" style={{ marginTop: 20 }}>
                  This is a warning alert — check it out!
                </Alert>
              )}
          </div>
        ) : null}
      </ContentHeader>
      <div className={classes.root}>
        <form noValidate autoComplete="off">
        <FormControl
              className={classes.margin}
              variant="outlined"
            >
              <InputLabel id="client-label">เครื่องรับชม</InputLabel>
              <Select
                labelId="client-label"
                id="client"
                value={clientID}
                onChange={clientIDhandleChange}
                style={{ width: 400 }}
              >
              {clients.map((item: EntClientEntity) => (
                  <MenuItem value={item.id}>{item.cLIENTNAME}</MenuItem>
                ))}
              </Select>
            </FormControl>

          <div>
            <FormControl
              className={classes.margin}
              variant="outlined"
            >
              <InputLabel id="user-label">สมาชิกห้องสมุด</InputLabel>
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
          </div>

          <FormControl
            className={classes.margin}
            variant="outlined"
          >
            <InputLabel id="bookingType">ประเภทของผู้ใช้งาน</InputLabel>
            <Select
              labelId="bookingType"
              id="bookingType"
              value={bookingtypeID}
              onChange={BookingIDhandleChange}
              style={{ width: 200 }}
            >
              {bookingtypes.map((item: EntBookingtype) => (
                <MenuItem value={item.id}>{item.bOOKTYPENAME}</MenuItem>
              ))}
            </Select>
          </FormControl>
          <div>
            <FormControl
              className={classes.margin}
              variant="outlined"
            >
              <TextField
                id="deathtime"
                label="ว/ด/ป เวลาเสียชีวิต"
                type="datetime-local"
                value={bookingdate}
                onChange={BookingDatehandleChange}
                className={classes.textField}
                InputLabelProps={{
                  shrink: true,
                }}
              />
            </FormControl>
          </div>

          <div className={classes.margin}>
            <Button
              onClick={() => {
                CreateBooking();
              }}
              variant="contained"
              color="primary"
            >
              [บันทึกการจอง]
           </Button>
            <Button
              style={{ marginLeft: 20 }}
              component={RouterLink}
              to="/welcome"
              variant="contained"
            >
              กลับ
           </Button>
          </div>
        </form>
      </div>
    </Content>
  </Page>
);
}