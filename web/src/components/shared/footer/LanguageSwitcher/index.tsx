import { Container, CSS } from "@nextui-org/react";
import { Language } from "store/i18n";
import { LanguageButton } from "../Language";

export type LanguageSwitcherProps = {
  languages: Language[];
};

export const LanguageSwitcher: React.FC<LanguageSwitcherProps> = ({
  languages,
}) => {
  return (
    <Container
      css={{ padding: "$10" }}
      display="flex"
      alignItems="center"
      justify="center"
    >
      {languages.map((x, key) => (
        <LanguageButton key={key} language={x} />
      ))}
    </Container>
  );
};
