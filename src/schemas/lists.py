from xml.etree.ElementTree import Element, SubElement

from pydantic import BaseModel


class CreateList(BaseModel):
    title: str


class UpdateList(BaseModel):
    title: str


class ListOut(BaseModel):
    title: str
    name: str
    web: str

    def to_opml(self) -> str:
        """Convert Podcast to OPML outline element"""
        outline = Element("outline", name=self.name, title=self.title, url=self.web)
        return outline

    def to_xml(self) -> Element:
        """Convert Podcast to XML format"""
        user_list = Element("User list")
        name = SubElement(user_list, "name")
        name.text = self.name
        title = SubElement(user_list, "title")
        title.text = self.title
        url = SubElement(user_list, "web")
        url.text = self.web

        return user_list
