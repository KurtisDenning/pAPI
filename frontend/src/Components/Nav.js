import React from "react";
import { Box, Flex, Heading } from "@chakra-ui/react";
const Nav = () => {
  return (
    <Box mx={"50px"} h={"10vh"}>
      <Flex h={"100%"}>
        <Heading mt={"auto"} as={"h2"} size={"2xl"}>
          pApi
        </Heading>
      </Flex>
    </Box>
  );
};

export default Nav;
