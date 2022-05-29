import React from "react";
import { Box, Divider, ListItem, UnorderedList } from "@chakra-ui/react";

const APICard = (props) => {
  return (
    <>
      <Box p={5}>
        <UnorderedList>
          <ListItem>
            <b>Title: </b>
            {props.name}
          </ListItem>

          <ListItem mt={5}>
            <b>Data 1: </b>
            {props.data1}
          </ListItem>

          <ListItem>
            <b>Data 2: </b>
            {props.data2}
          </ListItem>

          <ListItem>
            <b>Data 3: </b>
            {props.data3}
          </ListItem>

          <ListItem>
            <b>Data 4: </b>
            {props.data4}
          </ListItem>
        </UnorderedList>
      </Box>

      <Divider />
    </>
  );
};

export default APICard;
