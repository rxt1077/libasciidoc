package parser

//go:generate pigeon -optimize-parser -optimize-grammar -alternate-entrypoints DocumentRawLine,DocumentFragment,NormalGroup,AttributeDeclarationValueGroup,AttributeStructuredValue,DelimitedBlockElements,HeaderGroup,AttributeDeclarationValue,FileLocation,IncludedFileLine,MarkdownQuoteAttribution,BlockAttributes,InlineAttributes,TableColumnsAttribute,LineRanges,TagRanges,DocumentAuthorFullName -o parser.go parser.peg
