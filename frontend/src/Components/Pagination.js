import React from "react";
import { Box, Button } from "@chakra-ui/react";

const Pagination = ({ itemsPerPage, totalPosts, paginate }) => {
  const pageNumbers = [];

  for (let i = 1; i <= Math.ceil(totalPosts / itemsPerPage); i++) {
    pageNumbers.push(i);
  }

  return (
    <Box mt={5}>
      {pageNumbers.map((number) => (
        <Button
          variant={"link"}
          color={"#2222222"}
          // Dont have a way to set active yet
          // _active={{
          //   bg: "#222222",
          //   color: "#ffffff"
          // }}
          size={"md"}
          fontSize={["1.25rem", null, "1.5rem"]}
          ml={2} key={number}
          onClick={() => paginate(number)}>
          {number}
        </Button>
      ))}
    </Box>
  );
};

export default Pagination;
