import React from "react";
import { Box, Text, Flex } from "@chakra-ui/react";
const Nav = () => {
  return (
    <Box px={4}>
      <Flex h={"8vh"} alignItems={"center"}>
        <Text fontSize={"3xl"}>pAPI</Text>
      </Flex>
    </Box>
  );
};

export default Nav;
