import React from "react";
import { Link } from "react-router-dom";
import { Box, Flex, Heading } from "@chakra-ui/react";
const Nav = () => {
  return (
    <Box mx={"50px"} h={"10vh"}>
      <Flex h={"100%"}>
        <Heading fontSize={["2rem", "3rem"]} mt={"auto"} color={"#222222"}>
          <Link to={"/pAPI"}>pApi</Link>
        </Heading>
      </Flex>
    </Box>
  );
};

export default Nav;
