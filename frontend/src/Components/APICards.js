import React from "react";
import { useNavigate } from "react-router-dom";
import { Box, Divider, Text, Heading, Button, Flex } from "@chakra-ui/react";

const APICard = (props) => {
  let navigate = useNavigate();

  function nav() {
    navigate(`/api/${props._id}`);
    window.scrollTo({
      top: 0,
      behavior: "smooth",
    });
  }

  return (
    <>
      <Box p={5}>
        <Heading as="h4" size="md">
          {props.title}
        </Heading>
        <Text>{props.description}</Text>
      </Box>
      <Flex justify={"end"}>
        <Button
          my={5}
          onClick={() => {
            nav();
          }}
        >
          Explore
        </Button>
      </Flex>
      <Divider />
    </>
  );
};

export default APICard;
