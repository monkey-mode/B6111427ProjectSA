import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import { EntBooking } from '../../api/models/EntBooking';
const useStyles = makeStyles({
    table: {
        minWidth: 650,
    },
});

export default function ComponentsTable() {
    const classes = useStyles();
    const api = new DefaultApi();
    
    const [bookings, setBookings] = useState<EntBooking[]>(Array);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const getUsers = async () => {
            const res = await api.listBooking({ limit: 20, offset: 0 });
            setLoading(false);
            setBookings(res);
        };
        getUsers();
    }, [loading]);
    
    return (
        <TableContainer component={Paper}>
            <Table className={classes.table} aria-label="simple table">
                <TableHead>
                    <TableRow>
                        <TableCell align="center">หมายเลขการจอง</TableCell>
                        <TableCell align="center">ผู้จอง</TableCell>
                        <TableCell align="center">ประเภทการจอง</TableCell>
                        <TableCell align="center">เครื่องที่กำลังใช้งาน</TableCell>
                        <TableCell align="center">วันเวลาที่จอง</TableCell>
                        <TableCell align="center">เวลาที่เหลือ</TableCell>
                    </TableRow>
                </TableHead>
                <TableBody>
                    {bookings.map((item:any)  => (
                        <TableRow key={item.id}>
                            <TableCell align="center">{item.id}</TableCell>
                            <TableCell align="center">{item.edges.usedby.uSERNAME}</TableCell>
                            <TableCell align="center">{item.edges.book.bOOKTYPENAME}</TableCell>
                            <TableCell align="center">{item.edges.using.cLIENTNAME}</TableCell>
                            <TableCell align="center">{item.bOOKINGDATE}</TableCell>
                            <TableCell align="center">{item.tIMELEFT}</TableCell>
                        </TableRow>
                    ))}
                </TableBody>
            </Table>
        </TableContainer>
    );
}