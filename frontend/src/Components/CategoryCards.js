import { React } from "react";
import { Button, Text } from "@chakra-ui/react";

const CategoryCards = ({ setCategory, category, item }) => {
  function setState() {
    if (category === item) {
      setCategory(null);
    } else {
      setCategory(item);
    }
  }

  return (

    <Button
      fontSize={["1rem", "1.25rem", "1.5rem"]}
      size={["lg"]}
      mr={[6, 9, 12]}
      mb={[6, 9, 12]}
      variant={"link"}
      color={"#2222222"}
      _active={{
        bg: "#222222",
        color: "#ffffff"
      }}
      isActive={category === item}
      onClick={() => {
        setState();
      }}
    >
      <Text m={3}>
        {item.shortDesc}
      </Text>
    </Button>
  );
};

export default CategoryCards;
