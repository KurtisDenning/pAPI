import React from "react";
import { useNavigate } from "react-router-dom";
import { Box, Divider, Text, Button, Flex } from "@chakra-ui/react";

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
    <Box m={5}>
      <Box p={5}>
        <Text fontSize={["1.25rem", "1.5rem", "2rem"]}>
          {props.title}
        </Text>
        <Text noOfLines={1} fontSize={["1rem", null, "1.25rem"]} color="gray">{props.description}</Text>
      </Box>
      <Flex justify={"end"}>
        <Button
          fontSize={"1rem"}
          my={5}
          onClick={() => {
            nav();
          }}
        >
          Explore
        </Button>
      </Flex>
      <Divider />
    </Box>
  );
};

export default APICard;
