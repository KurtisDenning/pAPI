import React from "react";
import { Link } from "react-router-dom";
import { Box, Flex, Heading } from "@chakra-ui/react";

const Footer = () => {

  function scrollToTop() {
    window.scrollTo({
      top: 0,
      behavior: "smooth",
    })
  }

  return (
    <Box bg={"#222222"}>
      <Flex h={"33vh"} alignItems={"center"} justifyContent={"center"}>
        <Heading color={"white"} fontSize={["2rem", "3rem"]} >
          <Link onClick={scrollToTop} to={"/pAPI"}>pApi</Link>
        </Heading>
      </Flex>
    </Box>
  );
};

export default Footer;
