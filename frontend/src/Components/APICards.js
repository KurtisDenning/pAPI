import React from "react";
import { Box, Divider, Text, Heading } from "@chakra-ui/react";

const APICard = (props) => {
  return (
    <>
      <Box p={5}>
        <Heading as="h4" size="md">
          {props.title}
        </Heading>
        <Text>{props.description}</Text>
      </Box>
      <Divider />
    </>
  );
};

export default APICard;
