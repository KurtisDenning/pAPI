import { React } from "react";
import { Button } from "@chakra-ui/react";

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
      isActive={category === item}
      onClick={() => {
        setState();
      }}
    >
      {item.shortDesc}
    </Button>
  );
};

export default CategoryCards;
