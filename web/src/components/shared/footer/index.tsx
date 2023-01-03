import { styled } from "@nextui-org/react";
import { LanguageSwitcher } from "./LanguageSwitcher";

const Foot = styled("footer", {
  flexGrow: 1,
});

export const Footer: React.FC = () => {
  return (
    <Foot>
      <LanguageSwitcher languages={["en", "ru"]} />
    </Foot>
  );
};
